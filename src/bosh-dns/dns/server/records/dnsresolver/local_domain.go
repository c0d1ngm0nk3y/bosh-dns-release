package dnsresolver

import (
	"net"

	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/miekg/dns"
)

type LocalDomain struct {
	logger    logger.Logger
	logTag    string
	recordSet RecordSet
	shuffler  AnswerShuffler
}

//go:generate counterfeiter . AnswerShuffler

type AnswerShuffler interface {
	Shuffle(src []dns.RR) []dns.RR
}

//go:generate counterfeiter . RecordSet

type RecordSet interface {
	Resolve(domain string) ([]string, error)
}

func NewLocalDomain(logger logger.Logger, recordSet RecordSet, shuffler AnswerShuffler) LocalDomain {
	return LocalDomain{
		logger:    logger,
		logTag:    "LocalDomain",
		recordSet: recordSet,
		shuffler:  shuffler,
	}
}

func (d LocalDomain) Resolve(questionDomains []string, responseWriter dns.ResponseWriter, requestMsg *dns.Msg) *dns.Msg {
	answers, rCode := d.resolve(requestMsg.Question[0].Name, questionDomains)

	responseMsg := &dns.Msg{}
	responseMsg.Answer = answers
	responseMsg.SetRcode(requestMsg, rCode)
	responseMsg.Authoritative = true
	responseMsg.RecursionAvailable = false

	d.trimIfNeeded(responseWriter, responseMsg)

	return responseMsg
}

func (d LocalDomain) resolve(answerDomain string, questionDomains []string) ([]dns.RR, int) {
	answers := []dns.RR{}

	for _, questionDomain := range questionDomains {
		ips, err := d.recordSet.Resolve(questionDomain)
		if err != nil {
			d.logger.Error(d.logTag, "failed to get ip addresses: %v", err)
			return nil, dns.RcodeFormatError
		}

		for _, ip := range ips {
			answer := &dns.A{
				Hdr: dns.RR_Header{
					Name:   answerDomain,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    0,
				},
				A: net.ParseIP(ip),
			}

			answers = append(answers, answer)
		}
	}

	return d.shuffler.Shuffle(answers), dns.RcodeSuccess
}

func (LocalDomain) trimIfNeeded(responseWriter dns.ResponseWriter, resp *dns.Msg) {
	maxLength := dns.MaxMsgSize
	_, isUDP := responseWriter.RemoteAddr().(*net.UDPAddr)

	if isUDP {
		maxLength = 512
	}

	numAnswers := len(resp.Answer)

	for len(resp.Answer) > 0 && resp.Len() > maxLength {
		resp.Answer = resp.Answer[:len(resp.Answer)-1]
	}

	resp.Truncated = isUDP && len(resp.Answer) < numAnswers
}

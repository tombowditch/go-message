package messageutil

import (
	"testing"
)

var formatHeaderFieldTests = []struct {
	k, v      string
	formatted string
}{
	{
		k:         "From",
		v:         "Mitsuha Miyamizu <mitsuha.miyamizu@example.org>",
		formatted: "From: Mitsuha Miyamizu <mitsuha.miyamizu@example.org>\r\n",
	},
	{
		k:         "Subject",
		v:         "This is a very long subject, much longer than just the 76 characters limit that applies to message header fields",
		formatted: "Subject: This is a very long subject, much longer than just the 76\r\n characters limit that applies to message header fields\r\n",
	},
	{
		k:         "Subject",
		v:         "This is        yet          \t  another    subject          \t                   with many         whitespace      characters",
		formatted: "Subject: This is        yet          \t  another    subject          \t       \r\n            with many         whitespace      characters\r\n",
	},
	{
		k:         "DKIM-Signature",
		v:         "v=1;\r\n h=From:To:Reply-To:Subject:Message-ID:References:In-Reply-To:MIME-Version;\r\n d=example.org\r\n",
		formatted: "DKIM-Signature: v=1;\r\n h=From:To:Reply-To:Subject:Message-ID:References:In-Reply-To:MIME-Version;\r\n d=example.org\r\n",
	},
	{
		k:         "DKIM-Signature",
		v:         "v=1; h=From; d=example.org; b=AuUoFEfDxTDkHlLXSZEpZj79LICEps6eda7W3deTVFOk4yAUoqOB4nujc7YopdG5dWLSdNg6xNAZpOPr+kHxt1IrE+NahM6L/LbvaHutKVdkLLkpVaVVQPzeRDI009SO2Il5Lu7rDNH6mZckBdrIx0orEtZV4bmp/YzhwvcubU4=\r\n",
		formatted: "DKIM-Signature: v=1; h=From; d=example.org;\r\n b=AuUoFEfDxTDkHlLXSZEpZj79LICEps6eda7W3deTVFOk4yAUoqOB4nujc7YopdG5dWLSdNg6x\r\n NAZpOPr+kHxt1IrE+NahM6L/LbvaHutKVdkLLkpVaVVQPzeRDI009SO2Il5Lu7rDNH6mZckBdrI\r\n x0orEtZV4bmp/YzhwvcubU4=\r\n",
	},
}

func TestFormatHeaderField(t *testing.T) {
	for _, test := range formatHeaderFieldTests {
		s := FormatHeaderField(test.k, test.v)
		if s != test.formatted {
			t.Errorf("Expected formatted header to be %q but got %q", test.formatted, s)
		}
	}
}

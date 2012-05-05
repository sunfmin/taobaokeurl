package taobaokeurl

import (
	"testing"
)

func TestTaobaoUrl(t *testing.T) {
	_, numiid := ResolveTaobaokeUrl("http://s.click.taobao.com/t_8?e=7HZ6jHSTbIOSInHkOznE9rNtB%2Be6VNDPzWcbEoQbq5li2p9dLCcOL%2FehErnQaWCD%2BsGwrWZN%2Fo3LS5Wk4JtjPaTanY0GCD9Zj67z8F2980HCQg%3D%3D&p=mm_17142583_0_0&n=19")
	// TaobaoUrl("http://s.click.taobao.com/t_8?e=7HZ6jHSTbIWfpnAkjey47iKYlY7feCQjNK2KaUOsvRLZp60TqYTSBwrd9wPJYuFPwOgHbaucXNW%2F%2BBQBtT5NP9XaVbPQaZ33JVfI1hPGSWv7ryQ%3D&p=mm_17142583_0_0&n=19&ref=http%3A%2F%2Fwww.meilishuo.com%2Fu%2FEA_TzO%3Fcc%3Da")
	if numiid != "13466006969" {
		t.Error(numiid)
	}
}

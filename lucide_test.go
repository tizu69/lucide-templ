package lucide_test

import (
	"context"
	"io"
	"testing"

	"g.tizu.dev/lucide-templ"
)

type Component interface {
	Render(ctx context.Context, w io.Writer) error
}

// make sure lucide.Icon implements templ.Component
var _ Component = (*lucide.Icon)(nil)

func TestClass(t *testing.T) {
	expRaw := "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" ><path d=\"M5 12h14\" /><path d=\"M12 5v14\" /></svg>"
	expCls := "<svg class=\"foo bar\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" ><path d=\"M5 12h14\" /><path d=\"M12 5v14\" /></svg>"
	expEsc := "<svg class=\"&#34;foo&#34; &#34;bar&#34;\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" ><path d=\"M5 12h14\" /><path d=\"M12 5v14\" /></svg>"

	if gotRaw := lucide.Plus; string(gotRaw) != expRaw {
		t.Errorf("expected %q, got %q", expRaw, gotRaw)
	}
	if gotCls := lucide.Plus.C("foo bar"); string(gotCls) != expCls {
		t.Errorf("expected %q, got %q", expCls, gotCls)
	}
	if gotEsc := lucide.Plus.C("\"foo\" \"bar\""); string(gotEsc) != expEsc {
		t.Errorf("expected %q, got %q", expEsc, gotEsc)
	}
}

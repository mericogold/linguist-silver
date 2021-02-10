package api

import (
    "testing"

    "github.com/go-test/deep"
)

func TestRenewer_NewRenewer(t *testing.T) {
    t.Parallel()

    client, err := NewClient(DefaultConfig())
    if err != nil {
        t.Fatal(err)
    }

    cases := []struct {
        name string
        i    *RenewerInput
        e    *Renewer
        err  bool
    }{
        {
            "nil",
            nil,
            nil,
            true,
        },
        {
            "missing_secret",
            &RenewerInput{
                Secret: nil,
            },
            nil,
            true,
        },
        {
            "default_grace",
            &RenewerInput{
                Secret: &Secret{},
            },
            &Renewer{
                secret: &Secret{},
            },
            false,
        },
    }
}

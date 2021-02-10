func TestRenewer_NewRenewer(t *testing.T) {
    t.Parallel()

    client, err := NewClient(DefaultConfig())
    if err != nil {
        t.Fatal(err)
    }

    cases := []struct {
        name string
        err  bool
    }{
        {
            "nil",
        },
    }
}

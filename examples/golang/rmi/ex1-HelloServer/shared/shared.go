package shared

type Hello int

// All RPC calls in the default library must:
// Be called on a type
// Take a pointer to an argument
// Take a pointer to a reply
// Return an error
func (h *Hello) SayHello(args *string, reply *string) error {
	*reply = "Hello World"
	return nil
}

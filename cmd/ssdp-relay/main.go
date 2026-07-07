fmt.Println("Configured networks")
fmt.Println()

if err := network.Discover(cfg); err != nil {
	log.Fatal(err)
}

for _, network := range cfg.Networks {

	fmt.Printf("✓ %s\n", network.CIDR)

	if network.Interface != nil {
		fmt.Printf("    Interface : %s\n", network.Interface.Name)
	} else {
		fmt.Printf("    Interface : not found\n")
	}

	for _, addr := range network.Addresses {
		fmt.Printf("    Address   : %s\n", addr)
	}

	fmt.Println()
}
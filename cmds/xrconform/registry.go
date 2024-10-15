package main

func TestAll(td *TD) {
	td.DependsOn(TestXRegistry)
	td.Run(TestRegistry1)
	td.DependsOn(TestRegistry1a)
	td.DependsOn(TestRegistry2)
	td.DependsOn(TestRegistry2a)
	td.DependsOn(TestRegistry3)
	td.DependsOn(TestRegistry4)
	td.Run(TestRegistry5)
	td.DependsOn(TestRegistry6)
	td.DependsOn(TestRegistry6a)
	td.Run(TestRegistry7)
}

func TestRegistry0(td *TD) {
	td.DependsOn(TestServer)
	td.Log("testreg0 log msg")
}

func TestRegistry1(td *TD) {
	td.Log("tr1 info line")
	td.Warn("just a warning - 1")
}

func TestRegistry1a(td *TD) {
	/*
		td.Log("1omething happened1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		td.Log("2omething happened1234567890123 456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		td.Log("3omething happened12345678901234567890123456789012345678901234567890123456789012345 67890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		td.Log("4omething happened12345678901234567890123456789012345678901234567890123456789012345 67890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		td.Log("5omethinghappened12345678901234567890123456789012345678901234567890123456 78901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		td.Log("6omethinghappened1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	*/
	td.Log("checking valid names")
}

func TestRegistry2(td *TD) {
	td.Pass()
}

func TestRegistry2a(td *TD) {
	td.Log("something happened")
	td.Warn("A warning")
	td.Pass()
}

func TestRegistry3(td *TD) {
	td.Pass("subtest1")
}

func TestRegistry4(td *TD) {
	td.Pass("subtest2")
	td.Skip("Not implemented")
	td.Pass()
}

func TestRegistry5(td *TD) {
	td.Log("tr5 info line")
	td.Pass("subtest3")
	td.Fail("Not good")
}

func TestRegistry6(td *TD) {
	td.Fail("subtest4 asd a sda sd asd asd ads a ds asd asd a ds asd asd asd a sda d asd as da ds asd ")
}

func TestRegistry6a(td *TD) {
	td.Fail("subtest4 asd a sda sd asd asd ads a ds asd asd a ds asd asd asd a sda d asd as da ds asd ")
	td.Fail("xxx")
}

func TestRegistry7(td *TD) {
	td.Log("something happened")
	td.Log("a b c d e f g h i j k l m n o p q r s t u v w x y z 1 2 3 4 5 6 7 8 9 0 q w e r t y u u i o p a s d f g h j k l x c v b n m q w e r t y u i o p a s d f g h j j k l")
	td.Fail("subtest5")
	// td.Fail("subtest5 asd a sda sd asd asd ads a ds asd asd a ds asd asd asd a sda d asd as da ds asd ")
	td.Pass()
}

package closureFunction

type Gen struct {
}

func (g *Gen) MakeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

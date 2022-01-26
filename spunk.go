package main

type spunk struct{
}

func (s *spunk) makeShoe() iShoe{
	return &spunkShoe{
		shoe: shoe{
			logo:"spunk",
			size: 14,
		},
	}
}


func (s *spunk) makeShirt() iShirt{
	return &spunkShirt{
		shirt: shirt{
			logo:"spunk",
			size: 14,
		},
	}
}
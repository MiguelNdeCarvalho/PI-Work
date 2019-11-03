package foodchain

import "fmt"

// Verse returns a certain verse.
func Verse(n int) string {
	if n < 1 {
		return ""
	}
	chain := []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
	thing := chain[n-1]
	v := fmt.Sprintf("I know an old lady who swallowed a %s.\n", thing)
	if n == 2 {
		v += "It wriggled and jiggled and tickled inside her.\n"
	} else if n == 3 {
		v += "How absurd to swallow a bird!\n"
	} else if n == 4 {
		v += "Imagine that, to swallow a cat!\n"
	} else if n == 5 {
		v += "What a hog, to swallow a dog!\n"
	} else if n == 6 {
		v += "Just opened her throat and swallowed a goat!\n"
	} else if n == 7 {
		v += "I don't know how she swallowed a cow!\n"
	} else if n == 8 {
		v += "She's dead, of course!"
		return v
	}

	for i := n - 1; i > 2; i-- {
		v += fmt.Sprintf("She swallowed the %s to catch the %s.\n", chain[i], chain[i-1])
	}

	if n > 2 {
		v += "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n"
	}
	if n > 1 {
		v += "She swallowed the spider to catch the fly.\n"
	}
	v += "I don't know why she swallowed the fly. Perhaps she'll die."
	return v
}

// Verses returns certain verses.
func Verses(start, end int) string {
	vs := ""
	for i := start; i <= end; i++ {
		vs += fmt.Sprintf("%s\n\n", Verse(i))
	}
	return vs[0 : len(vs)-2]
}

// Song returns the whole song.
func Song() string {
	return Verses(1, 8)
}

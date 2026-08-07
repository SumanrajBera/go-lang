module externalPackage

go 1.26.5

replace github.com/SumanrajBera/reverseString v0.0.0 => ./reverseString

require (
	github.com/SumanrajBera/reverseString v0.0.0
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/google/uuid v1.6.0
)

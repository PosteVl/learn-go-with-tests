package iteration


func Repeat(sequence string, repeatCount int) string{
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += sequence
    }

    return repeated
}

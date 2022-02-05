package iteration

const repeatCount = 6

func Repeat(sequence string) string{
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += sequence
    }

    return repeated
}

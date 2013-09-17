= GotoPinYin

convert Simplified Chinese to Pinyin with GoLang

### install

    go get github.com/elvuel/gotopinyin

### Test

    go test

### Usage && Example

    package main

    import (
      "fmt"
      pinyin "github.com/elvuel/gotopinyin"
    )

    func main() {
      fmt.Printf("%s", pinyin.Convert("汉字", " "))
    }

### Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
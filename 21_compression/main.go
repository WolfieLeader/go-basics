package main

import "fmt"

const LOREM_IPSUM = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris finibus turpis tortor, et mollis tortor porttitor vel. Aenean gravida pretium risus, in tincidunt sem gravida vitae. Donec vulputate eros sapien, a hendrerit diam auctor ut. Ut lectus nulla, porttitor quis ex in, semper rutrum turpis. Proin a auctor arcu. Duis tortor lectus, tristique aliquam porta in, placerat id erat. Integer tempus dui eu augue ultrices bibendum. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Duis elementum metus in dui ultricies, ut faucibus ligula varius. Maecenas euismod vestibulum mi sed ornare. Sed ut velit et purus semper sollicitudin. Nulla cursus lorem et dolor porttitor, eget eleifend velit viverra. Mauris vel posuere tortor. Fusce dignissim, diam ac condimentum aliquam, sem arcu facilisis diam, sed porttitor enim lorem non ligula. Sed in tempus nunc, vel bibendum massa. Cras sed elit sit amet turpis pellentesque ultricies. Phasellus quis scelerisque justo, quis euismod est. Sed maximus tellus sit amet luctus blandit. Aenean lacinia felis arcu, vel feugiat quam ullamcorper eget. Cras felis orci, dignissim ac dapibus id, aliquet eu quam. Maecenas ut neque id lorem consectetur aliquam ac vitae tellus. Morbi urna tortor, tempus vel purus quis, pretium rhoncus magna. Integer sem velit, dapibus eget tempor id, aliquet vel nulla. Donec convallis fringilla purus eget fringilla. Integer in efficitur sapien. Suspendisse ullamcorper lacus ut nulla luctus, vitae varius odio venenatis. Aenean ligula erat, auctor finibus faucibus ut, tincidunt nec tortor. Mauris convallis leo et sagittis venenatis. Proin a nulla congue, vestibulum erat non, posuere eros. Nulla tempus sapien at est bibendum rhoncus. In lobortis porttitor consequat. Vestibulum imperdiet turpis dapibus cursus malesuada. Aenean convallis lorem nisi, eget convallis lorem imperdiet vitae. Vivamus mollis molestie felis. Nunc nec tortor id enim gravida porttitor et at mi. Fusce sed malesuada felis. Pellentesque eu turpis sit amet ex sodales consequat. Sed at lectus quis urna sagittis interdum. Integer hendrerit, arcu nec venenatis lobortis, purus arcu consequat est, vitae molestie magna sapien nec mi. Sed consequat tempor aliquam."

func main() {
	fmt.Println("Gzip Compress Example:")
	gzipExample()
	fmt.Println()

	fmt.Println("ZLib Compress Example:")
	zLibExample()
	fmt.Println()

	fmt.Println("Flate Compress Example:")
	flateExample()
	fmt.Println()
}

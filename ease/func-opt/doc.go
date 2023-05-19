// Functional options is a pattern in which you declare an opaque Option type
// that records information in some internal struct. You accept a variadic number
//
// of these Options and act upon the full information recorder by the options
// on the internal struct.
// https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis:w
// https://github.com/uber-go/guide/blob/master/style.md#functional-options
// https://coolshell.cn/articles/21146.html
package main

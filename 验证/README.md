1. 空切片==nil
通过F1和F2进行测试
2. []byte和[]uint8是同一个东西，通过反射.(type)匹配[]type,[]uint8都可，通过F3和F4进行测试
3. len获取实际长度，cap获取容量长度，对于channel也是如此，通过F5进行测试
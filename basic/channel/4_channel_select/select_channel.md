### channel与select

```golang
select{
    case <-chan1: 
    //如果chan1成功读到数据,则进行该case处理语句
    case chan2 <-1:
    //如果成功向chan2写入数据,则进行该case处理语句
    default:
    //如果上面都没有成功，则进入default处理流程
}
```


/**
 * 不可恢复的错误是检测到的错误，程序员无法处理它。 
 * 当发生这种错误时，panic!宏会被执行。panic!打印失败消息。 
 * 
 * ---------------------------------------
 * 
 * 两种情况:
 *      ->展开(Unwinding)：展开是一个清理它遇到的每个函数的堆栈内存中的数据的过程。 但是，展开过程需要大量工作。Unwinding的替代方案是Aborting。
 *      ->中止(Aborting)：中止是在不清除堆栈内存中的数据的情况下结束程序的过程。操作系统将删除数据。如果从展开切换到中止，那么需要添加以下语句：
 */

fn main(){
    panic!("?No such file exist?");

    //-----
    let v = vec![20,30,40];  
    print!("element of a vector is :",v[5]);  //Panic!!
}
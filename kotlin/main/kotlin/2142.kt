fun main(args: Array<String>) {
    var (a, b, c) = readLine()!!.split(' ').map(String::toLong)
    val (x, y, z) = readLine()!!.split(' ').map(String::toLong)
    var tmp: Long = 0

    if ( ((a+c)-x) < 0) {
        println("There are no miracles in life")
        return;
    } else {
        if (a>x) {
            tmp = a-x
        } else {
            c = (a + c) - x
        }
    }

    if ( ((b+c)-y) < 0) {
        println("There are no miracles in life")
        return;
    } else {
        if (b>y) {
            tmp += b-y
        } else {
            c = (b + c) - y
        }
    }

    if ((c+tmp)<z) {
        println("There are no miracles in life")
        return;
    }

    println("It is a kind of magic")
}
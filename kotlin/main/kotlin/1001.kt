fun main(args: Array<String>) {
    val nums = mutableListOf<Double>()
    while (true) {
        nums +=
            (readLine() ?: break)
                .split(' ')
                .filter { !it.isEmpty() }
                .map { it.toDouble() }
    }
    nums.reversed()
        .forEach { println(Math.sqrt(it)) }
}
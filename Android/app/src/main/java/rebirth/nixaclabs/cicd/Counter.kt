package rebirth.nixaclabs.cicd

class Counter {
    private var count: Int = 0

    fun getCount(): Int = count

    fun increment() {
        count++
    }

    fun decrement() {
        count--
    }

    fun reset() {
        count = 0
    }
}
package rebirth.nixaclabs.cicd

import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test

class CounterTest {
    private lateinit var counter: Counter

    @Before
    fun setup() {
        counter = Counter()
    }

    @Test
    fun testInitialCountIsZero() {
        assertEquals(0, counter.getCount())
    }

    @Test
    fun testIncrementWorks() {
        counter.increment()
        assertEquals(1, counter.getCount())
    }

    @Test
    fun testDecrementWorks() {
        counter.decrement()
        assertEquals(-1, counter.getCount())
    }

    @Test
    fun testMultipleIncrements() {
        counter.increment()
        counter.increment()
        counter.increment()
        assertEquals(3, counter.getCount())
    }

    @Test
    fun testMultipleDecrements() {
        counter.decrement()
        counter.decrement()
        assertEquals(-2, counter.getCount())
    }

    @Test
    fun testIncrementAndDecrement() {
        counter.increment()
        counter.increment()
        counter.decrement()
        assertEquals(1, counter.getCount())
    }
}
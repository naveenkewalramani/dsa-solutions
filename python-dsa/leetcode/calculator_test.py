import unittest
import calculator

class TestCalculator(unittest.TestCase):
    def test_addition(self):
        obj = calculator.Calculator
        self.assertEqual(obj.addition(5, 10), 15)
    
    def test_subtraction(self):
        obj = calculator.Calculator
        self.assertEqual(obj.subtraction(5, 10), -5)

    def test_multiply(self):
        obj = calculator.Calculator
        self.assertEqual(obj.multiply(5, 10), 50)

    def test_divison(self):
        obj = calculator.Calculator
        self.assertEqual(obj.divison(5, 10), 0.5)

    def test_power(self):
        obj = calculator.Calculator
        self.assertEqual(obj.power(5, 2), 25)

if __name__ == "__main__":
    unittest.main()


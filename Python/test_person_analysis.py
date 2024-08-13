import unittest
from person_analysis import person_analysis

class test_person_analysis(unittest.TestCase):

# Перевірка коректності розрахунку BMI.
    def test_calculate_bmi(self):
        person = person_analysis("John Doe", 30, 175, 70)
        self.assertAlmostEqual(person.calculate_bmi(), 22.86, places=2)

# Перевірка стану здоров'я "Underweight".
    def test_get_health_status_underweight(self):
        person = person_analysis("Jane Doe", 25, 165, 45)
        self.assertEqual(person.get_health_status(), "Underweight")

# Перевірка стану здоров'я "Normal weight".
    def test_get_health_status_normal_weight(self):
        person = person_analysis("John Doe", 30, 175, 70)
        self.assertEqual(person.get_health_status(), "Normal weight")

# Перевірка стану здоров'я "Overweight".
    def test_get_health_status_overweight(self):
        person = person_analysis("Jim Doe", 35, 180, 85)
        self.assertEqual(person.get_health_status(), "Overweight")

# Перевірка стану здоров'я "Obesity".
    def test_get_health_status_obesity(self):
        person = person_analysis("Jake Doe", 40, 170, 100)
        self.assertEqual(person.get_health_status(), "Obesity")
        
# Перевірка рядкового представлення об'єкта.
    def test_str(self):
        person = person_analysis("John Doe", 30, 175, 70)
        expected_string = ("Name: John Doe, Age: 30, Height: 175, Weight: 70, "
                           "BMI: 22.857142857142858, Health Status: Normal weight")
        self.assertEqual(str(person), expected_string)

if __name__ == '__main__':
    unittest.main()
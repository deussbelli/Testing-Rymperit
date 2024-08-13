class person_analysis:
    def __init__(self, name, age, height, weight):
        self.name = name
        self.age = age
        self.height = height
        self.weight = weight

    def calculate_bmi(self):
        return self.weight / ((self.height / 100) ** 2)

    def get_health_status(self):
        bmi = self.calculate_bmi()
        if bmi < 18.5:
            return "Underweight"
        elif 18.5 <= bmi < 24.9:
            return "Normal weight"
        elif 25 <= bmi < 29.9:
            return "Overweight"
        else:
            return "Obesity"

    def __str__(self):
        return (f"Name: {self.name}, Age: {self.age}, Height: {self.height}, Weight: {self.weight}, "
                f"BMI: {self.calculate_bmi()}, Health Status: {self.get_health_status()}")

<?php
use PHPUnit\Framework\TestCase;
require_once 'person_analysis.php';

class test_person_analysis extends TestCase {

# Перевірка коректності розрахунку BMI.
    public function testCalculateBMI() {
        $person = new person_analysis("John Doe", 30, 175, 70);
        $this->assertEquals(22.86, round($person->calculateBMI(), 2));
    }

# Перевірка стану здоров'я "Underweight".
    public function testGetHealthStatusUnderweight() {
        $person = new person_analysis("Jane Doe", 25, 165, 45);
        $this->assertEquals("Underweight", $person->getHealthStatus());
    }

# Перевірка стану здоров'я "Normal weight".
    public function testGetHealthStatusNormalWeight() {
        $person = new person_analysis("John Doe", 30, 175, 70);
        $this->assertEquals("Normal weight", $person->getHealthStatus());
    }

# Перевірка стану здоров'я "Overweight".
    public function testGetHealthStatusOverweight() {
        $person = new person_analysis("Jim Doe", 35, 180, 85);
        $this->assertEquals("Overweight", $person->getHealthStatus());
    }

# Перевірка стану здоров'я "Obesity".
    public function testGetHealthStatusObesity() {
        $person = new person_analysis("Jake Doe", 40, 170, 100);
        $this->assertEquals("Obesity", $person->getHealthStatus());
    }

# Перевірка рядкового представлення об'єкта.
    public function testToString() {
        $person = new person_analysis("John Doe", 30, 175, 70);
        $expectedString = "Name: John Doe, Age: 30, Height: 175, Weight: 70, BMI: 22.86, Health Status: Normal weight";
        $this->assertEquals($expectedString, (string)$person);
    }
}
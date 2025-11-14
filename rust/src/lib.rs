pub fn add(a: f64, b: f64) -> f64 {
    a + b
}

pub fn sub(a: f64, b: f64) -> f64 {
    a - b
}

pub fn mul(a: f64, b: f64) -> f64 {
    a * b
}

pub fn div(a: f64, b: f64) -> f64 {
    if b == 0.0 {
        panic!("division by zero");
    }
    a / b
}

#[cfg(test)]
mod tests {
    use super::*;
    use rand::Rng;

    #[test]
    fn test_add() {
        // ЗАМЕНИТЬ ЭТОТ БЛОК КОДА:
        assert_eq!(add(2.0, 3.0), 5.0);
        assert_eq!(add(-2.0, 3.0), 1.0);
        assert_eq!(add(0.5, 0.5), 1.0);
    }

    #[test]
    fn test_sub() {
        let mut rng = rand::thread_rng();
        let a: f64 = rng.gen();
        let b: f64 = rng.gen();
        assert_eq!(sub(a, b), a - b);
    }

    #[test]
    fn test_mul() {
        let mut rng = rand::thread_rng();
        let a: f64 = rng.gen();
        let b: f64 = rng.gen();
        assert_eq!(mul(a, b), a * b);
    }

    #[test]
    fn test_div() {
        let mut rng = rand::thread_rng();
        let a: f64 = rng.gen();
        let b: f64 = rng.gen();
        if b != 0.0 {
            assert_eq!(div(a, b), a / b);
        }
    }

    #[test]
    #[should_panic(expected = "division by zero")]
    fn test_div_zero() {
        div(1.0, 0.0);
    }
}
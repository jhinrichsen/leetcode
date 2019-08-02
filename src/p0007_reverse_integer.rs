fn reverse_integer(x: i32) -> i32 {
    // nothing to do for single digits
    if x > -10 && x < 10 {
        return x;
    }

    let sign = x.signum();
    let mut n = x.abs();

    let mut y: i32 = 0;
    loop {
        let n10 = n / 10;
        let last_digit = n - (n10 * 10);
        // handle overflow
        match y.checked_mul(10) {
            Some(val) => y = val,
            None => return 0,
        }
        match y.checked_add(last_digit) {
            Some(val) => y = val,
            None => return 0,
        }
        if n10 == 0 {
            break;
        }
        n = n10;
    }
    sign * y
}

#[cfg(test)]
mod tests {
    #[test]
    fn reverse_integer_0() {
        assert_eq!(0, super::reverse_integer(0));
    }

    #[test]
    fn reverse_integer_3() {
        assert_eq!(3, super::reverse_integer(3));
    }

    #[test]
    fn reverse_integer_minus_3() {
        assert_eq!(-3, super::reverse_integer(-3));
    }

    #[test]
    fn reverse_integer_123() {
        assert_eq!(321, super::reverse_integer(123));
    }

    #[test]
    fn reverse_integer_minus_123() {
        assert_eq!(-321, super::reverse_integer(-123));
    }

    #[test]
    fn reverse_integer_120() {
        assert_eq!(21, super::reverse_integer(120));
    }

    #[test]
    fn reverse_integer_1534236469() {
        assert_eq!(0, super::reverse_integer(1534236469));
    }
}

// try-into is unstable as of current Rust 1.31 in LeetCode
// use std::convert::TryInto;

// returning 0 for errors is non-idiomatic in Rust
// using str as variable name is non-idiomatic in Rust
// using String instead of str is non-idiomatic in Rust
fn my_atoi(str: String) -> i32 {
    let radix = 10;
    let mut sign: i32 = 1;
    let mut pass = 1;
    let mut n: i32 = 0;
    for c in str.chars() {
        if pass == 1 {
            if c == ' ' {
                continue;
            }
            pass = 2;
        }
        if pass == 2 {
            if c == '+' {
                sign = 1;
                pass = 3;
                continue;
            }
            if c == '-' {
                sign = -1;
                pass = 3;
                continue;
            }
            if c.is_digit(radix) {
                pass = 3;
            } else {
                return 0;
            }
        }
        if pass == 3 {
            if c.is_digit(radix) {
                n = n.saturating_mul(10);
                // let m: i32 = c.to_digit(radix).unwrap().try_into().unwrap();
                let m: i32 = c.to_digit(radix).unwrap() as i32;
                n = n.saturating_add(m * sign);
            } else {
                pass = 4;
            }
        }
        if pass == 4 {
            // overread any trailing non-digit characters
            break;
        }
    }
    n
}

#[cfg(test)]
mod tests {
    #[test]
    fn my_atoi_42() {
        assert_eq!(42, super::my_atoi("42".to_string()));
    }

    #[test]
    fn my_atoi_spaces_minus_42() {
        assert_eq!(-42, super::my_atoi("   -42".to_string()));
    }

    #[test]
    fn my_atoi_4193_with_words() {
        assert_eq!(4193, super::my_atoi("4193 with words".to_string()));
    }

    #[test]
    fn my_atoi_words_and_987() {
        assert_eq!(0, super::my_atoi("words and 987".to_string()));
    }

    #[test]
    fn my_atoi_underflow() {
        assert_eq!(-2147483648, super::my_atoi("-91283472332".to_string()));
    }

    #[test]
    fn my_atoi_plus_minus() {
        assert_eq!(0, super::my_atoi("+-2".to_string()));
    }
}

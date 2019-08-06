pub fn check_possibility(nums: Vec<i32>) -> bool {
    let is_monotone = |v: &[i32]| -> bool {
        for i in 0..v.len() - 1 {
            if v[i] > v[i + 1] {
                return false;
            }
        }
        true
    };

    let mut bak;
    let mut clone = nums.to_vec();
    for i in 0..nums.len() {
        let src_index = {
            if i == 0 {
                1
            } else {
                i - 1
            }
        };
        bak = clone[i];
        clone[i] = clone[src_index];
        if is_monotone(&clone) {
            return true;
        }
        clone[i] = bak;
    }
    false
}

#[cfg(test)]
mod tests {
    #[test]
    fn check_possiblity_423() {
        assert_eq!(true, super::check_possibility(vec![4, 2, 3]));
    }

    #[test]
    fn check_possibility_421() {
        assert_eq!(false, super::check_possibility(vec![4, 2, 1]));
    }

    #[test]
    fn check_possibility_3423() {
        assert_eq!(false, super::check_possibility(vec![3, 4, 2, 3]));
    }

    #[test]
    fn check_possibility_111() {
        assert_eq!(true, super::check_possibility(vec![1, 1, 1]));
    }

    #[test]
    fn check_possibility_32324() {
        assert_eq!(false, super::check_possibility(vec![3, 2, 3, 2, 4]));
    }

    #[test]
    fn check_possibility_23324() {
        assert_eq!(true, super::check_possibility(vec![2, 3, 3, 2, 4]));
    }

    #[test]
    fn check_possibility_132() {
        assert_eq!(true, super::check_possibility(vec![1, 3, 2]));
    }
}

pub fn check_possibility(nums: Vec<i32>) -> bool {
    let mut anomalies = 0;
    for i in 1..nums.len() {
        if nums[i - 1] <= nums[i] {
            continue;
        }

        anomalies += 1;
        if anomalies > 1 {
            return false;
        }
        let has_right = i + 1 < nums.len();
        if has_right {
            let monotone = nums[i - 1] <= nums[i + 1];
            if !monotone {
                return false;
            }
        }
    }
    true
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

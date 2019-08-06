pub fn check_possibility(nums: Vec<i32>) -> bool {
    let anomalies = |ns: &[i32]| {
        let mut n = 0;
        for i in 0..ns.len() - 1 {
            if ns[i] > ns[i + 1] {
                n += 1;
            }
        }
        n
    };
    if anomalies(&nums[..]) == 0 {
        return true;
    }
    if anomalies(&nums[..]) > 1 {
        return false;
    }
    let idx = |ns: &[i32]| {
        for i in 0..ns.len() - 1 {
            if ns[i] > ns[i + 1] {
                return i;
            }
        }
        ns.len()
    };
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

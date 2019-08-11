pub fn third_max(nums: Vec<i32>) -> i32 {
	let mut res = nums.to_owned();
	res.sort_unstable();
	res.dedup();
	let idx = if res.len() < 3 {
		res.len() - 1
	} else {
		res.len() - 3
	};
	res[idx]
}

#[cfg(test)]
mod tests {
	#[test]
	fn third_max_3_2_1() {
		assert_eq!(1, super::third_max(vec![3, 2, 1]));
	}

	#[test]
	fn third_max_1_2() {
		assert_eq!(2, super::third_max(vec![1, 2]));
	}

	#[test]
	fn third_max_2_2_3_1() {
		assert_eq!(1, super::third_max(vec![2, 2, 3, 1]));
	}
}

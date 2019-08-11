// String as parameter is not idiomatic Rust
pub fn defang(address: String) -> String {
	let mut s = String::new();
	for c in address.chars() {
		if c == '.' {
			s.push_str("[.]");
		} else {
			s.push(c);
		}
	}
	s
}

#[cfg(test)]
mod tests {
	#[test]
	fn defang_1111() {
		assert_eq!("1[.]1[.]1[.]1", super::defang("1.1.1.1".to_string()));
	}

	#[test]
	fn defang_ip() {
		assert_eq!(
			"192[.]168[.]178[.]1",
			super::defang("192.168.178.1".to_string())
		);
	}
}

# Sort Inside
매우 간단한 정렬 문제이다.
한줄로 입력된 자연수를 각 자리수마다 처리하면 되기 때문에
간단히 문자열로 해석해서 풀었다.

# 문제 해결 아이디어
1. 한줄로 주어지는 input을 string으로 저장한다.
2. string을 loop를 통해 rune으로 풀어준다.
3. 풀린 rune값을 int처리해서 미리 만들어둔 카운팅 배열에 횟수를 저장한다.
4. 배열에 저장된 횟수만큼 역으로 풀어서 출력한다.

# 시간 복잡도
N
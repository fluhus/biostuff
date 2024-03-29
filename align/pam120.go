package align

func init() {
	PAM120 = SubstitutionMatrix{
		{'A', 'A'}: 3,
		{'A', 'B'}: 0,
		{'A', 'C'}: -3,
		{'A', 'D'}: 0,
		{'A', 'E'}: 0,
		{'A', 'F'}: -4,
		{'A', 'G'}: 1,
		{'A', 'H'}: -3,
		{'A', 'I'}: -1,
		{'A', 'K'}: -2,
		{'A', 'L'}: -3,
		{'A', 'M'}: -2,
		{'A', 'N'}: -1,
		{'A', 'P'}: 1,
		{'A', 'Q'}: -1,
		{'A', 'R'}: -3,
		{'A', 'S'}: 1,
		{'A', 'T'}: 1,
		{'A', 'V'}: 0,
		{'A', 'W'}: -7,
		{'A', 'X'}: -1,
		{'A', 'Y'}: -4,
		{'A', 'Z'}: -1,
		{'A', Gap}: -8,
		{'B', 'A'}: 0,
		{'B', 'B'}: 4,
		{'B', 'C'}: -6,
		{'B', 'D'}: 4,
		{'B', 'E'}: 3,
		{'B', 'F'}: -5,
		{'B', 'G'}: 0,
		{'B', 'H'}: 1,
		{'B', 'I'}: -3,
		{'B', 'K'}: 0,
		{'B', 'L'}: -4,
		{'B', 'M'}: -4,
		{'B', 'N'}: 3,
		{'B', 'P'}: -2,
		{'B', 'Q'}: 0,
		{'B', 'R'}: -2,
		{'B', 'S'}: 0,
		{'B', 'T'}: 0,
		{'B', 'V'}: -3,
		{'B', 'W'}: -6,
		{'B', 'X'}: -1,
		{'B', 'Y'}: -3,
		{'B', 'Z'}: 2,
		{'B', Gap}: -8,
		{'C', 'A'}: -3,
		{'C', 'B'}: -6,
		{'C', 'C'}: 9,
		{'C', 'D'}: -7,
		{'C', 'E'}: -7,
		{'C', 'F'}: -6,
		{'C', 'G'}: -4,
		{'C', 'H'}: -4,
		{'C', 'I'}: -3,
		{'C', 'K'}: -7,
		{'C', 'L'}: -7,
		{'C', 'M'}: -6,
		{'C', 'N'}: -5,
		{'C', 'P'}: -4,
		{'C', 'Q'}: -7,
		{'C', 'R'}: -4,
		{'C', 'S'}: 0,
		{'C', 'T'}: -3,
		{'C', 'V'}: -3,
		{'C', 'W'}: -8,
		{'C', 'X'}: -4,
		{'C', 'Y'}: -1,
		{'C', 'Z'}: -7,
		{'C', Gap}: -8,
		{'D', 'A'}: 0,
		{'D', 'B'}: 4,
		{'D', 'C'}: -7,
		{'D', 'D'}: 5,
		{'D', 'E'}: 3,
		{'D', 'F'}: -7,
		{'D', 'G'}: 0,
		{'D', 'H'}: 0,
		{'D', 'I'}: -3,
		{'D', 'K'}: -1,
		{'D', 'L'}: -5,
		{'D', 'M'}: -4,
		{'D', 'N'}: 2,
		{'D', 'P'}: -3,
		{'D', 'Q'}: 1,
		{'D', 'R'}: -3,
		{'D', 'S'}: 0,
		{'D', 'T'}: -1,
		{'D', 'V'}: -3,
		{'D', 'W'}: -8,
		{'D', 'X'}: -2,
		{'D', 'Y'}: -5,
		{'D', 'Z'}: 3,
		{'D', Gap}: -8,
		{'E', 'A'}: 0,
		{'E', 'B'}: 3,
		{'E', 'C'}: -7,
		{'E', 'D'}: 3,
		{'E', 'E'}: 5,
		{'E', 'F'}: -7,
		{'E', 'G'}: -1,
		{'E', 'H'}: -1,
		{'E', 'I'}: -3,
		{'E', 'K'}: -1,
		{'E', 'L'}: -4,
		{'E', 'M'}: -3,
		{'E', 'N'}: 1,
		{'E', 'P'}: -2,
		{'E', 'Q'}: 2,
		{'E', 'R'}: -3,
		{'E', 'S'}: -1,
		{'E', 'T'}: -2,
		{'E', 'V'}: -3,
		{'E', 'W'}: -8,
		{'E', 'X'}: -1,
		{'E', 'Y'}: -5,
		{'E', 'Z'}: 4,
		{'E', Gap}: -8,
		{'F', 'A'}: -4,
		{'F', 'B'}: -5,
		{'F', 'C'}: -6,
		{'F', 'D'}: -7,
		{'F', 'E'}: -7,
		{'F', 'F'}: 8,
		{'F', 'G'}: -5,
		{'F', 'H'}: -3,
		{'F', 'I'}: 0,
		{'F', 'K'}: -7,
		{'F', 'L'}: 0,
		{'F', 'M'}: -1,
		{'F', 'N'}: -4,
		{'F', 'P'}: -5,
		{'F', 'Q'}: -6,
		{'F', 'R'}: -5,
		{'F', 'S'}: -3,
		{'F', 'T'}: -4,
		{'F', 'V'}: -3,
		{'F', 'W'}: -1,
		{'F', 'X'}: -3,
		{'F', 'Y'}: 4,
		{'F', 'Z'}: -6,
		{'F', Gap}: -8,
		{'G', 'A'}: 1,
		{'G', 'B'}: 0,
		{'G', 'C'}: -4,
		{'G', 'D'}: 0,
		{'G', 'E'}: -1,
		{'G', 'F'}: -5,
		{'G', 'G'}: 5,
		{'G', 'H'}: -4,
		{'G', 'I'}: -4,
		{'G', 'K'}: -3,
		{'G', 'L'}: -5,
		{'G', 'M'}: -4,
		{'G', 'N'}: 0,
		{'G', 'P'}: -2,
		{'G', 'Q'}: -3,
		{'G', 'R'}: -4,
		{'G', 'S'}: 1,
		{'G', 'T'}: -1,
		{'G', 'V'}: -2,
		{'G', 'W'}: -8,
		{'G', 'X'}: -2,
		{'G', 'Y'}: -6,
		{'G', 'Z'}: -2,
		{'G', Gap}: -8,
		{'H', 'A'}: -3,
		{'H', 'B'}: 1,
		{'H', 'C'}: -4,
		{'H', 'D'}: 0,
		{'H', 'E'}: -1,
		{'H', 'F'}: -3,
		{'H', 'G'}: -4,
		{'H', 'H'}: 7,
		{'H', 'I'}: -4,
		{'H', 'K'}: -2,
		{'H', 'L'}: -3,
		{'H', 'M'}: -4,
		{'H', 'N'}: 2,
		{'H', 'P'}: -1,
		{'H', 'Q'}: 3,
		{'H', 'R'}: 1,
		{'H', 'S'}: -2,
		{'H', 'T'}: -3,
		{'H', 'V'}: -3,
		{'H', 'W'}: -3,
		{'H', 'X'}: -2,
		{'H', 'Y'}: -1,
		{'H', 'Z'}: 1,
		{'H', Gap}: -8,
		{'I', 'A'}: -1,
		{'I', 'B'}: -3,
		{'I', 'C'}: -3,
		{'I', 'D'}: -3,
		{'I', 'E'}: -3,
		{'I', 'F'}: 0,
		{'I', 'G'}: -4,
		{'I', 'H'}: -4,
		{'I', 'I'}: 6,
		{'I', 'K'}: -3,
		{'I', 'L'}: 1,
		{'I', 'M'}: 1,
		{'I', 'N'}: -2,
		{'I', 'P'}: -3,
		{'I', 'Q'}: -3,
		{'I', 'R'}: -2,
		{'I', 'S'}: -2,
		{'I', 'T'}: 0,
		{'I', 'V'}: 3,
		{'I', 'W'}: -6,
		{'I', 'X'}: -1,
		{'I', 'Y'}: -2,
		{'I', 'Z'}: -3,
		{'I', Gap}: -8,
		{'K', 'A'}: -2,
		{'K', 'B'}: 0,
		{'K', 'C'}: -7,
		{'K', 'D'}: -1,
		{'K', 'E'}: -1,
		{'K', 'F'}: -7,
		{'K', 'G'}: -3,
		{'K', 'H'}: -2,
		{'K', 'I'}: -3,
		{'K', 'K'}: 5,
		{'K', 'L'}: -4,
		{'K', 'M'}: 0,
		{'K', 'N'}: 1,
		{'K', 'P'}: -2,
		{'K', 'Q'}: 0,
		{'K', 'R'}: 2,
		{'K', 'S'}: -1,
		{'K', 'T'}: -1,
		{'K', 'V'}: -4,
		{'K', 'W'}: -5,
		{'K', 'X'}: -2,
		{'K', 'Y'}: -5,
		{'K', 'Z'}: -1,
		{'K', Gap}: -8,
		{'L', 'A'}: -3,
		{'L', 'B'}: -4,
		{'L', 'C'}: -7,
		{'L', 'D'}: -5,
		{'L', 'E'}: -4,
		{'L', 'F'}: 0,
		{'L', 'G'}: -5,
		{'L', 'H'}: -3,
		{'L', 'I'}: 1,
		{'L', 'K'}: -4,
		{'L', 'L'}: 5,
		{'L', 'M'}: 3,
		{'L', 'N'}: -4,
		{'L', 'P'}: -3,
		{'L', 'Q'}: -2,
		{'L', 'R'}: -4,
		{'L', 'S'}: -4,
		{'L', 'T'}: -3,
		{'L', 'V'}: 1,
		{'L', 'W'}: -3,
		{'L', 'X'}: -2,
		{'L', 'Y'}: -2,
		{'L', 'Z'}: -3,
		{'L', Gap}: -8,
		{'M', 'A'}: -2,
		{'M', 'B'}: -4,
		{'M', 'C'}: -6,
		{'M', 'D'}: -4,
		{'M', 'E'}: -3,
		{'M', 'F'}: -1,
		{'M', 'G'}: -4,
		{'M', 'H'}: -4,
		{'M', 'I'}: 1,
		{'M', 'K'}: 0,
		{'M', 'L'}: 3,
		{'M', 'M'}: 8,
		{'M', 'N'}: -3,
		{'M', 'P'}: -3,
		{'M', 'Q'}: -1,
		{'M', 'R'}: -1,
		{'M', 'S'}: -2,
		{'M', 'T'}: -1,
		{'M', 'V'}: 1,
		{'M', 'W'}: -6,
		{'M', 'X'}: -2,
		{'M', 'Y'}: -4,
		{'M', 'Z'}: -2,
		{'M', Gap}: -8,
		{'N', 'A'}: -1,
		{'N', 'B'}: 3,
		{'N', 'C'}: -5,
		{'N', 'D'}: 2,
		{'N', 'E'}: 1,
		{'N', 'F'}: -4,
		{'N', 'G'}: 0,
		{'N', 'H'}: 2,
		{'N', 'I'}: -2,
		{'N', 'K'}: 1,
		{'N', 'L'}: -4,
		{'N', 'M'}: -3,
		{'N', 'N'}: 4,
		{'N', 'P'}: -2,
		{'N', 'Q'}: 0,
		{'N', 'R'}: -1,
		{'N', 'S'}: 1,
		{'N', 'T'}: 0,
		{'N', 'V'}: -3,
		{'N', 'W'}: -4,
		{'N', 'X'}: -1,
		{'N', 'Y'}: -2,
		{'N', 'Z'}: 0,
		{'N', Gap}: -8,
		{'P', 'A'}: 1,
		{'P', 'B'}: -2,
		{'P', 'C'}: -4,
		{'P', 'D'}: -3,
		{'P', 'E'}: -2,
		{'P', 'F'}: -5,
		{'P', 'G'}: -2,
		{'P', 'H'}: -1,
		{'P', 'I'}: -3,
		{'P', 'K'}: -2,
		{'P', 'L'}: -3,
		{'P', 'M'}: -3,
		{'P', 'N'}: -2,
		{'P', 'P'}: 6,
		{'P', 'Q'}: 0,
		{'P', 'R'}: -1,
		{'P', 'S'}: 1,
		{'P', 'T'}: -1,
		{'P', 'V'}: -2,
		{'P', 'W'}: -7,
		{'P', 'X'}: -2,
		{'P', 'Y'}: -6,
		{'P', 'Z'}: -1,
		{'P', Gap}: -8,
		{'Q', 'A'}: -1,
		{'Q', 'B'}: 0,
		{'Q', 'C'}: -7,
		{'Q', 'D'}: 1,
		{'Q', 'E'}: 2,
		{'Q', 'F'}: -6,
		{'Q', 'G'}: -3,
		{'Q', 'H'}: 3,
		{'Q', 'I'}: -3,
		{'Q', 'K'}: 0,
		{'Q', 'L'}: -2,
		{'Q', 'M'}: -1,
		{'Q', 'N'}: 0,
		{'Q', 'P'}: 0,
		{'Q', 'Q'}: 6,
		{'Q', 'R'}: 1,
		{'Q', 'S'}: -2,
		{'Q', 'T'}: -2,
		{'Q', 'V'}: -3,
		{'Q', 'W'}: -6,
		{'Q', 'X'}: -1,
		{'Q', 'Y'}: -5,
		{'Q', 'Z'}: 4,
		{'Q', Gap}: -8,
		{'R', 'A'}: -3,
		{'R', 'B'}: -2,
		{'R', 'C'}: -4,
		{'R', 'D'}: -3,
		{'R', 'E'}: -3,
		{'R', 'F'}: -5,
		{'R', 'G'}: -4,
		{'R', 'H'}: 1,
		{'R', 'I'}: -2,
		{'R', 'K'}: 2,
		{'R', 'L'}: -4,
		{'R', 'M'}: -1,
		{'R', 'N'}: -1,
		{'R', 'P'}: -1,
		{'R', 'Q'}: 1,
		{'R', 'R'}: 6,
		{'R', 'S'}: -1,
		{'R', 'T'}: -2,
		{'R', 'V'}: -3,
		{'R', 'W'}: 1,
		{'R', 'X'}: -2,
		{'R', 'Y'}: -5,
		{'R', 'Z'}: -1,
		{'R', Gap}: -8,
		{'S', 'A'}: 1,
		{'S', 'B'}: 0,
		{'S', 'C'}: 0,
		{'S', 'D'}: 0,
		{'S', 'E'}: -1,
		{'S', 'F'}: -3,
		{'S', 'G'}: 1,
		{'S', 'H'}: -2,
		{'S', 'I'}: -2,
		{'S', 'K'}: -1,
		{'S', 'L'}: -4,
		{'S', 'M'}: -2,
		{'S', 'N'}: 1,
		{'S', 'P'}: 1,
		{'S', 'Q'}: -2,
		{'S', 'R'}: -1,
		{'S', 'S'}: 3,
		{'S', 'T'}: 2,
		{'S', 'V'}: -2,
		{'S', 'W'}: -2,
		{'S', 'X'}: -1,
		{'S', 'Y'}: -3,
		{'S', 'Z'}: -1,
		{'S', Gap}: -8,
		{'T', 'A'}: 1,
		{'T', 'B'}: 0,
		{'T', 'C'}: -3,
		{'T', 'D'}: -1,
		{'T', 'E'}: -2,
		{'T', 'F'}: -4,
		{'T', 'G'}: -1,
		{'T', 'H'}: -3,
		{'T', 'I'}: 0,
		{'T', 'K'}: -1,
		{'T', 'L'}: -3,
		{'T', 'M'}: -1,
		{'T', 'N'}: 0,
		{'T', 'P'}: -1,
		{'T', 'Q'}: -2,
		{'T', 'R'}: -2,
		{'T', 'S'}: 2,
		{'T', 'T'}: 4,
		{'T', 'V'}: 0,
		{'T', 'W'}: -6,
		{'T', 'X'}: -1,
		{'T', 'Y'}: -3,
		{'T', 'Z'}: -2,
		{'T', Gap}: -8,
		{'V', 'A'}: 0,
		{'V', 'B'}: -3,
		{'V', 'C'}: -3,
		{'V', 'D'}: -3,
		{'V', 'E'}: -3,
		{'V', 'F'}: -3,
		{'V', 'G'}: -2,
		{'V', 'H'}: -3,
		{'V', 'I'}: 3,
		{'V', 'K'}: -4,
		{'V', 'L'}: 1,
		{'V', 'M'}: 1,
		{'V', 'N'}: -3,
		{'V', 'P'}: -2,
		{'V', 'Q'}: -3,
		{'V', 'R'}: -3,
		{'V', 'S'}: -2,
		{'V', 'T'}: 0,
		{'V', 'V'}: 5,
		{'V', 'W'}: -8,
		{'V', 'X'}: -1,
		{'V', 'Y'}: -3,
		{'V', 'Z'}: -3,
		{'V', Gap}: -8,
		{'W', 'A'}: -7,
		{'W', 'B'}: -6,
		{'W', 'C'}: -8,
		{'W', 'D'}: -8,
		{'W', 'E'}: -8,
		{'W', 'F'}: -1,
		{'W', 'G'}: -8,
		{'W', 'H'}: -3,
		{'W', 'I'}: -6,
		{'W', 'K'}: -5,
		{'W', 'L'}: -3,
		{'W', 'M'}: -6,
		{'W', 'N'}: -4,
		{'W', 'P'}: -7,
		{'W', 'Q'}: -6,
		{'W', 'R'}: 1,
		{'W', 'S'}: -2,
		{'W', 'T'}: -6,
		{'W', 'V'}: -8,
		{'W', 'W'}: 12,
		{'W', 'X'}: -5,
		{'W', 'Y'}: -2,
		{'W', 'Z'}: -7,
		{'W', Gap}: -8,
		{'X', 'A'}: -1,
		{'X', 'B'}: -1,
		{'X', 'C'}: -4,
		{'X', 'D'}: -2,
		{'X', 'E'}: -1,
		{'X', 'F'}: -3,
		{'X', 'G'}: -2,
		{'X', 'H'}: -2,
		{'X', 'I'}: -1,
		{'X', 'K'}: -2,
		{'X', 'L'}: -2,
		{'X', 'M'}: -2,
		{'X', 'N'}: -1,
		{'X', 'P'}: -2,
		{'X', 'Q'}: -1,
		{'X', 'R'}: -2,
		{'X', 'S'}: -1,
		{'X', 'T'}: -1,
		{'X', 'V'}: -1,
		{'X', 'W'}: -5,
		{'X', 'X'}: -2,
		{'X', 'Y'}: -3,
		{'X', 'Z'}: -1,
		{'X', Gap}: -8,
		{'Y', 'A'}: -4,
		{'Y', 'B'}: -3,
		{'Y', 'C'}: -1,
		{'Y', 'D'}: -5,
		{'Y', 'E'}: -5,
		{'Y', 'F'}: 4,
		{'Y', 'G'}: -6,
		{'Y', 'H'}: -1,
		{'Y', 'I'}: -2,
		{'Y', 'K'}: -5,
		{'Y', 'L'}: -2,
		{'Y', 'M'}: -4,
		{'Y', 'N'}: -2,
		{'Y', 'P'}: -6,
		{'Y', 'Q'}: -5,
		{'Y', 'R'}: -5,
		{'Y', 'S'}: -3,
		{'Y', 'T'}: -3,
		{'Y', 'V'}: -3,
		{'Y', 'W'}: -2,
		{'Y', 'X'}: -3,
		{'Y', 'Y'}: 8,
		{'Y', 'Z'}: -5,
		{'Y', Gap}: -8,
		{'Z', 'A'}: -1,
		{'Z', 'B'}: 2,
		{'Z', 'C'}: -7,
		{'Z', 'D'}: 3,
		{'Z', 'E'}: 4,
		{'Z', 'F'}: -6,
		{'Z', 'G'}: -2,
		{'Z', 'H'}: 1,
		{'Z', 'I'}: -3,
		{'Z', 'K'}: -1,
		{'Z', 'L'}: -3,
		{'Z', 'M'}: -2,
		{'Z', 'N'}: 0,
		{'Z', 'P'}: -1,
		{'Z', 'Q'}: 4,
		{'Z', 'R'}: -1,
		{'Z', 'S'}: -1,
		{'Z', 'T'}: -2,
		{'Z', 'V'}: -3,
		{'Z', 'W'}: -7,
		{'Z', 'X'}: -1,
		{'Z', 'Y'}: -5,
		{'Z', 'Z'}: 4,
		{'Z', Gap}: -8,
		{Gap, 'A'}: -8,
		{Gap, 'B'}: -8,
		{Gap, 'C'}: -8,
		{Gap, 'D'}: -8,
		{Gap, 'E'}: -8,
		{Gap, 'F'}: -8,
		{Gap, 'G'}: -8,
		{Gap, 'H'}: -8,
		{Gap, 'I'}: -8,
		{Gap, 'K'}: -8,
		{Gap, 'L'}: -8,
		{Gap, 'M'}: -8,
		{Gap, 'N'}: -8,
		{Gap, 'P'}: -8,
		{Gap, 'Q'}: -8,
		{Gap, 'R'}: -8,
		{Gap, 'S'}: -8,
		{Gap, 'T'}: -8,
		{Gap, 'V'}: -8,
		{Gap, 'W'}: -8,
		{Gap, 'X'}: -8,
		{Gap, 'Y'}: -8,
		{Gap, 'Z'}: -8,
		{Gap, Gap}: 0,
	}
}

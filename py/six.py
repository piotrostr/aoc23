class Solution:
    def __init__(self, input_path: str):
        with open(input_path, "r+") as f:
            self.input_contents = f.readlines()

    def _parse_input(self) -> tuple[list[str], list[str]]:
        assert len(self.input_contents) == 2
        timestr, diststr = self.input_contents
        times = self._parse_line(timestr)
        distances = self._parse_line(diststr)
        return times, distances

    def _parse_line(self, fileline: str) -> list[int]:
        line = fileline.replace("\n", "")
        return [int(item) for item in line.split(" ") if item.isnumeric()]

    def main(self):
        """
        this can be optimized by a lot
        """
        times, distances = self._parse_input()
        number_of_ways_product = 1
        for time, min_distance in zip(times, distances):
            ways_to_complete = 0
            for k in range(time):
                diff = time - k
                distance = diff * k
                if distance > min_distance:
                    ways_to_complete += 1
            number_of_ways_product *= ways_to_complete

        time = int("".join(str(time) for time in times))
        min_distance = int("".join(str(dist) for dist in distances))
        start, end = -1, -1
        for k in range(time):
            diff = time - k
            distance = diff * k
            if distance > min_distance:
                start = k
                break
        for k in range(time - 1, -1, -1):
            diff = time - k
            distance = diff * k
            if distance > min_distance:
                end = k
                break
        print(end - start + 1)


if __name__ == "__main__":
    Solution("six.txt").main()

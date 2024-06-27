defmodule Day1 do
  def parse_file(pid) do
    stream = IO.stream(pid, :line)

    lines = Enum.map(stream, fn x -> String.trim(x) end)

    numbers = Enum.map(lines, fn line ->
      Enum.filter(String.graphemes(line), fn char ->
        case Integer.parse(char) do
          :error -> false
          _ -> true
        end
      end)
    end)

    total = Enum.reduce(numbers, 0, fn list, acc ->
      str = List.first(list) <> List.last(list)
      {int, ""} = Integer.parse(str)

      acc + int
    end)

    IO.puts(total)
  end

  def show_error_and_quit(error) do
    case error do
      :no_file ->
        IO.puts("No file argument")
      :enoent ->
        IO.puts("File does not exist")
      _ ->
        IO.puts("Could not load file")
    end

    IO.puts("Usage: day1.exs <file>")

    Kernel.exit({:shutdown, 1})
  end

  def parse_command_line() do
    case List.first(System.argv) do
      nil -> show_error_and_quit(:no_file)
      path -> path
    end
  end
end

path = Day1.parse_command_line()

case File.open(path, [:read, :utf8]) do
  {:ok, pid} -> Day1.parse_file(pid)
  {:error, error} -> Day1.show_error_and_quit(error)
end

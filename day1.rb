def char_is_integer?(char)
  char.to_i.to_s == char
end

def parse_first_and_last_integer(integer_chars)
  "#{integer_chars.first}#{integer_chars.last}".to_i
end

begin
  path = ARGF.filename
  lines = File.readlines(path)

  stripped_lines = lines.map(&:strip)

  integer_chars = stripped_lines.map do |line|
    line.chars.select do |char|
      char_is_integer?(char)
    end
  end

  numbers = integer_chars.map do |array|
    parse_first_and_last_integer(array)
  end

  puts numbers.inject(:+)
rescue
  puts "could not open file: #{argf.filename}"
  puts 'Usage: day1.rb <file>'
end

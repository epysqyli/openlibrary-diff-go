require 'smarter_csv'
require 'parallel'
require 'json'

headers = %w[type key revision last_modified json]
input_file = 'old.txt'
output_file = 'old_keys.txt'

File.open(output_file, 'w') do |f|
  SmarterCSV.process(input_file,
                     chunk_size: 50_000,
                     col_sep: "\t",
                     user_provided_headers: headers,
                     quote_char: "\x00",
                     invalid_byte_sequence: '') do |chunk|
    Parallel.each(chunk) do |row|
      next if row.nil?

      key = row[:key]&.split('/')&.last
      f.write "#{key}\n"
    end
  end
  file.close
end

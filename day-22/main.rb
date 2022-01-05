# frozen_string_literal: true

class App
  def run
    input = File.readlines("./data/input.txt", chomp: true)
  end
end

App.new.run

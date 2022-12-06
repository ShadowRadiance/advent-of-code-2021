# frozen_string_literal: true

require "./reactor_core"

class App
  def run
    input = File.readlines("./data/input.txt", chomp: true)

    core = ReactorCore.new(input)
    core.reboot(init_only: true)
    puts core.on_count
  end
end

App.new.run

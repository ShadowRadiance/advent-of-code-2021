class App
  def initialize
  end

  def run
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run

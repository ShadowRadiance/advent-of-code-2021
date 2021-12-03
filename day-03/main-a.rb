require "./diagnostic_report"

class App
  def initialize
    @report = DiagnosticReport.new(load_data_from_input)
  end

  def run
    puts @report.gamma_rate * @report.epsilon_rate
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
  end
end

App.new.run

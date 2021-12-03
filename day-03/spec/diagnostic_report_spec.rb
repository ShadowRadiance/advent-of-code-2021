require "./diagnostic_report"

RSpec.describe 'Diagnostic Report' do
  subject { DiagnosticReport.new(data) }
  before do
  end

  context 'with no data' do
    let(:data) { "" }

    it "reports the correct gamma rate" do
      expect(subject.gamma_rate).to eq(0)
    end
    it "reports the correct epsilon rate" do
      expect(subject.epsilon_rate).to eq(0)
    end
  end

  context 'with the sample data' do
    let(:data) {
      %w[
        00100
        11110
        10110
        10111
        10101
        01111
        00111
        11100
        10000
        11001
        00010
        01010
      ]
    }

    it "reports the correct gamma rate" do
      expect(subject.gamma_rate).to eq(22)
    end
    it "reports the correct epsilon rate" do
      expect(subject.epsilon_rate).to eq(9)
    end
  end
end
require "./depth_scanner"

RSpec.describe 'Depth Measurement' do
  subject { DepthScanner.new(data) }

  context 'with no data' do
    let(:data) { [] }

    it "reports the correct number of depth increases" do
      expect(subject.increases).to eq(0)
    end

  end

  context 'with the sample data' do
    let(:data) {
      [
        199, # first value - not an increase or decrease
        200, # increase
        208, # increase
        210, # increase
        200, # decrease
        207, # increase
        240, # increase
        269, # increase
        260, # decrease
        263, # increase
      ]
    }

    it "reports the correct number of depth increases" do
      expect(subject.increases).to eq(7)
    end
  end

end
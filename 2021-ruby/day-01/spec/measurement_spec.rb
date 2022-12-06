require "./depth_scanner"

RSpec.describe 'Depth Measurement' do
  subject { DepthScanner.new(data, strategy: strategy) }
  let(:strategy) { SimpleDepthScanner.new }

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

    context 'with a sliding_window scanner' do
      let(:strategy) { SlidingWindowScanner.new(3) }

      it "reports the correct number of depth increases" do
        expect(subject.increases).to eq(5)
      end
    end
  end

end
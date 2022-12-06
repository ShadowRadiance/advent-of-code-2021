require "./submarine"

RSpec.describe 'Submarine' do
  subject { Submarine.new }
  before do
    subject.follow_instructions(data.split("\n"))
  end

  context 'with no data' do
    let(:data) { "" }

    it "reports the correct location" do
      expect(subject.x_pos).to eq(0)
    end
    it "reports the correct depth" do
      expect(subject.depth).to eq(0)
    end
  end

  context 'with the sample data' do
    let(:data) {
      <<~DATA
        forward 5
        down 5
        forward 8
        up 3
        down 8
        forward 2
      DATA
    }

    it "reports the correct location" do
      # expect(subject.x_pos).to eq(15)
      expect(subject.x_pos).to eq(15)
    end
    it "reports the correct depth" do
      # expect(subject.depth).to eq(10)
      expect(subject.depth).to eq(60)
    end
  end
end
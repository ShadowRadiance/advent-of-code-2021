require './crab_optimizer'

RSpec.describe "Crab Optimizer" do
  subject  { CrabOptimizer.new(positions) }

  let(:positions) {
    [16,1,2,0,4,2,7,1,2,14]
  }

  describe "#optimal_alignment_fuel" do
    it "should find the minimum value" do
      # expect(subject.optimal_alignment_fuel).to eq(37)
      expect(subject.optimal_alignment_fuel).to eq(168)
    end
  end

  describe "#optimal_alignment" do
    it "should find the cheapest alignment" do
      # expect(subject.optimal_alignment).to eq(2)
      expect(subject.optimal_alignment).to eq(5)
    end
  end
end
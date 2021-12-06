require './lanternfish_simulator'

RSpec.describe "Lanternfish Simulator" do
  let(:simulator)  { LanternfishSimulator.new(data) }
  let(:data) {
    [3,4,3,1,2]
  }
  before do
    simulator.run(days: days)
  end
  
  describe "#all" do
    subject { simulator.all }

    context "for 1 day" do
      let(:days) { 1 }
      it { is_expected.to eq([2,3,2,0,1])}
    end

    context "for 11 day" do
      let(:days) { 11 }
      it { is_expected.to eq([6,0,6,4,5,6,0,1,1,2,6,7,8,8,8])}
    end
  end
  
  describe "#count" do
    subject { simulator.count }

    context "for 1 day" do
      let(:days) { 1 }
      it { is_expected.to eq(5)}
    end

    context "for 11 day" do
      let(:days) { 11 }
      it { is_expected.to eq(15)}
    end

    context "for 80 days" do
      let(:days) { 80 }
      it { is_expected.to eq(5934) }
    end
  end
end
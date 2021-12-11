require "./octopus"

RSpec.describe Octopus do
  subject { described_class.new(energy) }

  describe "#power_up" do

    context "when an octopus has 0 energy" do
      let(:energy) { 0 }

      before do
        subject.power_up
      end

      it { is_expected.to have_attributes(energy_level: 1) }
    end

    context "when an octopus has 8 energy" do
      let(:energy) { 8 }

      before do
        subject.power_up
      end
    
      it { is_expected.to have_attributes(energy_level: 9) }
    end  

    context "when an octopus has 9 energy" do
      let(:energy) { 9 }

      before do
        subject.power_up
      end

      it { is_expected.to have_attributes(energy_level: 10) }
    end

    context "when an octopus has 200 energy" do
      let(:energy) { 200 }

      before do
        subject.power_up
      end

      it { is_expected.to have_attributes(energy_level: 201) }
    end
  end

  describe "#flash_cascade" do
    let(:energy) { 10 }

    context "without high powered neighbors" do
      let(:neighbors) { [Octopus.new(1), Octopus.new(2), Octopus.new(3)] }

      before do
        subject.neighbors = neighbors
        neighbors.each do |neighbor|
          expect(neighbor).to receive(:flash_cascade).and_call_original
          expect(neighbor).to receive(:power_up).and_call_original
        end
  
        subject.flash_cascade(initiator: true)
      end
  
      it { is_expected.to have_flashed }
      it { is_expected.to have_attributes(energy_level: 10) }

      it "affects the neighbors" do
        expect(neighbors[0].energy_level).to eq(2)
        expect(neighbors[1].energy_level).to eq(3)
        expect(neighbors[2].energy_level).to eq(4)
      end
    end    
  end

  describe "#reset" do
    let(:energy) { 10 }
    
    before do
      subject.reset
    end

    it { is_expected.to have_attributes(energy_level: 0) }
    it { is_expected.not_to have_flashed }
  end

end

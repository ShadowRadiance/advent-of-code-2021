require "./game"

RSpec.describe Game do
  
  subject { Game.new(data) }

  context "with the simple data" do
    let(:data) {
      [
        [1,1,1,1,1],
        [1,9,9,9,1],
        [1,9,1,9,1],
        [1,9,9,9,1],
        [1,1,1,1,1],
      ]
    }
    before do
      steps.times { subject.step }
    end

    context "after 1 step" do
      let(:steps) { 1 }
      it { is_expected.to have_attributes(energy_board: [
        [3,4,5,4,3],
        [4,0,0,0,4],
        [5,0,0,0,5],
        [4,0,0,0,4],
        [3,4,5,4,3],
      ])}
    end

    context "after 2 step" do
      let(:steps) { 2 }
      it { is_expected.to have_attributes(energy_board: [
        [4,5,6,5,4],
        [5,1,1,1,5],
        [6,1,1,1,6],
        [5,1,1,1,5],
        [4,5,6,5,4],
      ])}
    end
  end

  context "with the complex data" do
    let(:data) {
      [
        [5,4,8,3,1,4,3,2,2,3],
        [2,7,4,5,8,5,4,7,1,1],
        [5,2,6,4,5,5,6,1,7,3],
        [6,1,4,1,3,3,6,1,4,6],
        [6,3,5,7,3,8,5,4,7,8],
        [4,1,6,7,5,2,4,6,4,5],
        [2,1,7,6,8,4,1,7,2,1],
        [6,8,8,2,8,8,1,1,3,4],
        [4,8,4,6,8,4,8,5,5,4],
        [5,2,8,3,7,5,1,5,2,6],
      ]
    }

    before do
      steps.times { subject.step }
    end

    context "after 1 steps" do
      let(:steps) { 1 }

      it "has the correct board state" do
        expect(subject.total_flashes).to eq(0)
        expect(subject.energy_board).to eq([
          [6,5,9,4,2,5,4,3,3,4],
          [3,8,5,6,9,6,5,8,2,2],
          [6,3,7,5,6,6,7,2,8,4],
          [7,2,5,2,4,4,7,2,5,7],
          [7,4,6,8,4,9,6,5,8,9],
          [5,2,7,8,6,3,5,7,5,6],
          [3,2,8,7,9,5,2,8,3,2],
          [7,9,9,3,9,9,2,2,4,5],
          [5,9,5,7,9,5,9,6,6,5],
          [6,3,9,4,8,6,2,6,3,7],
        ])
      end
    end

    context "after 10 steps" do
      let(:steps) { 10 }
      it "has the correct board state" do
        expect(subject.total_flashes).to eq(204)
        expect(subject.energy_board).to eq([
          [0,4,8,1,1,1,2,9,7,6],
          [0,0,3,1,1,1,2,0,0,9],
          [0,0,4,1,1,1,2,5,0,4],
          [0,0,8,1,1,1,1,4,0,6],
          [0,0,9,9,1,1,1,3,0,6],
          [0,0,9,3,5,1,1,2,3,3],
          [0,4,4,2,3,6,1,1,3,0],
          [5,5,3,2,2,5,2,3,5,0],
          [0,5,3,2,2,5,0,6,0,0],
          [0,0,3,2,2,4,0,0,0,0],
        ])
      end
    end

    context "after 100 steps" do
      let(:steps) { 100 }
      it "has the correct board state" do
        expect(subject.total_flashes).to eq(1656)
        expect(subject.energy_board).to eq([
          [0,3,9,7,6,6,6,8,6,6],
          [0,7,4,9,7,6,6,9,1,8],
          [0,0,5,3,9,7,6,9,3,3],
          [0,0,0,4,2,9,7,8,2,2],
          [0,0,0,4,2,2,9,8,9,2],
          [0,0,5,3,2,2,2,8,7,7],
          [0,5,3,2,2,2,2,9,6,6],
          [9,3,2,2,2,2,8,9,6,6],
          [7,9,2,2,2,8,6,8,6,6],
          [6,7,8,9,9,9,8,7,6,6],
        ])
      end
    end
  end
end

require './height_map'

RSpec.describe "Height Map" do
  describe ".new" do
    let(:data) {
      [
        [1,2,3],
        [4,5,6],
        [7,8,9],
      ]
    }
    subject { HeightMap.new(data) }
    
    it "loads the data into the height map" do
      expect(subject.data.first).to eq([1,2,3])
      expect(subject.data.last).to eq([7,8,9])
    end
  end

  context "with data loaded from a string" do
    let(:data) {
      <<~DATA
      2199943210
      3987894921
      9856789892
      8767896789
      9899965678
      DATA
    }
    subject { HeightMap.load_string(data) }
    
    describe ".load_string" do
      it "loads the data into the height map" do
        expect(subject.data.first).to eq([2,1,9,9,9,4,3,2,1,0])
        expect(subject.data.last).to eq([9,8,9,9,9,6,5,6,7,8])
      end
    end

    describe "#risk_level_at_point" do
      it "determines the risk level at a locations" do
        expect(subject.risk_level_at_point(1,0)).to eq(2)
        expect(subject.risk_level_at_point(6,4)).to eq(6)
      end
    end

    describe "#low_points" do
      it "determines the low point locations" do
        expect(subject.low_points).to eq([[1,0], [9,0], [2,2], [6,4]])
      end
    end
  
    describe "#low_point_risk_levels" do
      it "determines the risk levels of the low points" do
        expect(subject.low_point_risk_levels).to eq([2, 1, 6, 6])
      end
    end
  end

end
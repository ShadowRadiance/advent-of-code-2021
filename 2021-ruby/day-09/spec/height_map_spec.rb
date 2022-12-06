require './height_map'

RSpec.describe "Height Map" do
  def pt(x,y); HeightMap::Point.new(x,y); end

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

    describe "#risk_level_at" do
      it "determines the risk level at a locations" do
        expect(subject.risk_level_at(pt(1,0))).to eq(2)
        expect(subject.risk_level_at(pt(6,4))).to eq(6)
      end
    end

    describe "#low_points" do
      it "determines the low point locations" do
        expect(subject.low_points).to eq([
          pt(1,0), 
          pt(9,0),
          pt(2,2),
          pt(6,4)])
      end
    end
  
    describe "#low_point_risk_levels" do
      it "determines the risk levels of the low points" do
        expect(subject.low_point_risk_levels).to eq([2, 1, 6, 6])
      end
    end

    describe "#basins" do
      it "determines the basins" do
        expect(subject.basins).to have_attributes(size: 4)
        expect(subject.basins.last.members).to eq([
          pt(0,0),
          pt(1,0), 
          pt(0,1),
        ])
        expect(subject.basins.first.members).to eq([
                            pt(2,1), pt(3,1), pt(4,1),
                   pt(1,2), pt(2,2), pt(3,2), pt(4,2), pt(5,2),
          pt(0,3), pt(1,3), pt(2,3), pt(3,3), pt(4,3),
                   pt(1,4),
        ])
      end
    end
  end

end
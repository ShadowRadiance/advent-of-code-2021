# frozen_string_literal: true
require './vent/line'
require './vent/point'

RSpec.describe "Vent Line" do
  let(:line) { Vent::Line.new(start_point, end_point) }

  describe "#points" do
    subject { line.points }

    context "with a horizontal line" do
      let(:start_point) { Vent::Point.new(0,9) }
      let(:end_point) { Vent::Point.new(5,9) }
        
      it { is_expected.to be_an Array }
      it { is_expected.to all (be_a Vent::Point) }
      it { is_expected.to include(Vent::Point.new(0,9)) }
      it { is_expected.to include(Vent::Point.new(1,9)) }
      it { is_expected.to include(Vent::Point.new(2,9)) }
      it { is_expected.to include(Vent::Point.new(3,9)) }
      it { is_expected.to include(Vent::Point.new(4,9)) }
      it { is_expected.to include(Vent::Point.new(5,9)) }
    end

    context "with a BACKWARDS horizontal line" do
      let(:start_point) { Vent::Point.new(5,9) }
      let(:end_point) { Vent::Point.new(0,9) }
        
      it { is_expected.to be_an Array }
      it { is_expected.to all (be_a Vent::Point) }
      it { is_expected.to include(Vent::Point.new(0,9)) }
      it { is_expected.to include(Vent::Point.new(1,9)) }
      it { is_expected.to include(Vent::Point.new(2,9)) }
      it { is_expected.to include(Vent::Point.new(3,9)) }
      it { is_expected.to include(Vent::Point.new(4,9)) }
      it { is_expected.to include(Vent::Point.new(5,9)) }
    end

    context "with a vertical line" do
      let(:start_point) { Vent::Point.new(9,0) }
      let(:end_point) { Vent::Point.new(9,5) }
        
      it { is_expected.to be_an Array }
      it { is_expected.to all (be_a Vent::Point) }
      it { is_expected.to include(Vent::Point.new(9,0)) }
      it { is_expected.to include(Vent::Point.new(9,1)) }
      it { is_expected.to include(Vent::Point.new(9,2)) }
      it { is_expected.to include(Vent::Point.new(9,3)) }
      it { is_expected.to include(Vent::Point.new(9,4)) }
      it { is_expected.to include(Vent::Point.new(9,5)) }
    end

    context "with a BACKWARDS vertical line" do
      let(:start_point) { Vent::Point.new(9,5) }
      let(:end_point) { Vent::Point.new(9,0) }
        
      it { is_expected.to be_an Array }
      it { is_expected.to all (be_a Vent::Point) }
      it { is_expected.to include(Vent::Point.new(9,0)) }
      it { is_expected.to include(Vent::Point.new(9,1)) }
      it { is_expected.to include(Vent::Point.new(9,2)) }
      it { is_expected.to include(Vent::Point.new(9,3)) }
      it { is_expected.to include(Vent::Point.new(9,4)) }
      it { is_expected.to include(Vent::Point.new(9,5)) }
    end

  end

  # context "with a diagonal line" do
  #   let(:start_point) { Vent::Point.new(0,0) }
  #   let(:end_point) { Vent::Point.new(3,3) }

  #   it { is_expected.to be_an Array }
  #   it { is_expected.to all (be_a Vent::Point) }
  #   it { is_expected.to include(Vent::Point.new(0,0)) }
  #   it { is_expected.to include(Vent::Point.new(1,1)) }
  #   it { is_expected.to include(Vent::Point.new(2,2)) }
  #   it { is_expected.to include(Vent::Point.new(3,3)) }
  # end
end
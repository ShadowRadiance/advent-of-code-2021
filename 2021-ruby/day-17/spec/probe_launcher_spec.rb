# frozen_string_literal: true

require './probe_launcher'

RSpec.describe ProbeLauncher do
  subject { ProbeLauncher.new(target) }

  context "with the example target" do
    let(:target) { [20..30, -10..-5] }

    context "in the beginning" do
      it "has not hit or missed any shots" do
        expect(subject.shots_taken).to eq(0)
        expect(subject.shots_on_target.size).to eq(0)
      end
    end

    context "take a shot with v 7,2" do
      before { subject.shoot(7,2) }

      it "has hit one shot and missed none" do
        expect(subject.shots_taken).to eq(1)
        expect(subject.shots_on_target.size).to eq(1)
        expect(subject.shots_on_target.first).to eq({ initial_velocity: Vector[7,2],
                                                      height: 3,
                                                      final_location: Vector[28,-7] })
      end
    end

    context "take a shot with v 6,3" do
      before { subject.shoot(6,3) }

      it "has hit one shot and missed none" do
        expect(subject.shots_taken).to eq(1)
        expect(subject.shots_on_target.size).to eq(1)
        expect(subject.shots_on_target.first).to eq({ initial_velocity: Vector[6,3],
                                                      height: 6,
                                                      final_location: Vector[21,-9] })
      end
    end

    context "take a shot with v 9,0" do
      before { subject.shoot(9,0) }

      it "has hit one shot and missed none" do
        expect(subject.shots_taken).to eq(1)
        expect(subject.shots_on_target.size).to eq(1)
        expect(subject.shots_on_target.first).to eq({ initial_velocity: Vector[9,0],
                                                      height: 0,
                                                      final_location: Vector[30,-6] })
      end
    end

    context "take a shot with v 17,-4" do
      before { subject.shoot(17,-4) }

      it "has hit no shots and missed one" do
        expect(subject.shots_taken).to eq(1)
        expect(subject.shots_on_target.size).to eq(0)
        expect(subject.shots_on_target.first).to be_nil
      end
    end

    context "take a bunch of shots" do
      before do
        subject.shoot(7,2)
        subject.shoot(6,3)
        subject.shoot(9,0)
        subject.shoot(6,9)
        subject.shoot(17,-4)
      end

      it "has hit four shots and missed one" do
        expect(subject.shots_taken).to eq(5)
        expect(subject.shots_on_target.size).to eq(4)
        expect(subject.highest_accurate_shot).to eq({ initial_velocity: Vector[6,9],
                                                      height: 45,
                                                      final_location: Vector[21, -10] })
        expect(subject.puzzle).to eq(45)
      end
    end

    describe "determine the highest possible shot" do
      before { subject.determine_highest_possible_shot }

      it "has determined the highest possible shot" do
        expect(subject.highest_accurate_shot).to eq({ initial_velocity: Vector[6,9],
                                                      height: 45,
                                                      final_location: Vector[21, -10] })
        expect(subject.puzzle).to eq(45)
      end
    end

    describe "determine total possible shots" do
      before { subject.determine_all_possible_shots }

      it "determines all possible shots" do
        expect(subject.shots_on_target.size).to eq(112)

        expectations = [
          Vector[23,-10],  Vector[25,-9],   Vector[27,-5],   Vector[29,-6],   Vector[22,-6],
          Vector[21,-7],   Vector[9,0],     Vector[27,-7],   Vector[24,-5],   Vector[25,-7],
          Vector[26,-6],   Vector[25,-5],   Vector[6,8],     Vector[11,-2],   Vector[20,-5],
          Vector[29,-10],  Vector[6,3],     Vector[28,-7],   Vector[8,0],     Vector[30,-6],
          Vector[29,-8],   Vector[20,-10],  Vector[6,7],     Vector[6,4],     Vector[6,1],
          Vector[14,-4],   Vector[21,-6],   Vector[26,-10],  Vector[7,-1],    Vector[7,7],
          Vector[8,-1],    Vector[21,-9],   Vector[6,2],     Vector[20,-7],   Vector[30,-10],
          Vector[14,-3],   Vector[20,-8],   Vector[13,-2],   Vector[7,3],     Vector[28,-8],
          Vector[29,-9],   Vector[15,-3],   Vector[22,-5],   Vector[26,-8],   Vector[25,-8],
          Vector[25,-6],   Vector[15,-4],   Vector[9,-2],    Vector[15,-2],   Vector[12,-2],
          Vector[28,-9],   Vector[12,-3],   Vector[24,-6],   Vector[23,-7],   Vector[25,-10],
          Vector[7,8],     Vector[11,-3],   Vector[26,-7],   Vector[7,1],     Vector[23,-9],
          Vector[6,0],     Vector[22,-10],  Vector[27,-6],   Vector[8,1],     Vector[22,-8],
          Vector[13,-4],   Vector[7,6],     Vector[28,-6],   Vector[11,-4],   Vector[12,-4],
          Vector[26,-9],   Vector[7,4],     Vector[24,-10],  Vector[23,-8],   Vector[30,-8],
          Vector[7,0],     Vector[9,-1],    Vector[10,-1],   Vector[26,-5],   Vector[22,-9],
          Vector[6,5],     Vector[7,5],     Vector[23,-6],   Vector[28,-10],  Vector[10,-2],
          Vector[11,-1],   Vector[20,-9],   Vector[14,-2],   Vector[29,-7],   Vector[13,-3],
          Vector[23,-5],   Vector[24,-8],   Vector[27,-9],   Vector[30,-7],   Vector[28,-5],
          Vector[21,-10],  Vector[7,9],     Vector[6,6],     Vector[21,-5],   Vector[27,-10],
          Vector[7,2],     Vector[30,-9],   Vector[21,-8],   Vector[22,-7],   Vector[24,-9],
          Vector[20,-6],   Vector[6,9],     Vector[29,-5],   Vector[8,-2],    Vector[27,-8],
          Vector[30,-5],   Vector[24,-7]
        ]
        expectations.each do |expectation|
          expect(subject.shots_on_target.map {|shot| shot[:initial_velocity] }).to include(expectation)
        end
      end
    end
  end
end

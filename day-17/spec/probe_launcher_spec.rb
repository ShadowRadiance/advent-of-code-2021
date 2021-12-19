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
  end
end
